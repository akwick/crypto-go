# Acknowledgments

## Images

* *7TH_BIRTHDAY.png* - Artwork by [Ashley Willis](https://twitter.com/ashleymcnamara?ref_src=twsrc%5Egoogle%7Ctwcamp%5Eserp%7Ctwgr%5Eauthor) licensed under 
[CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/) found on [GitHub](https://github.com/ashleymcnamara/gophers). 

## Creating the ECB Image

I followed the tutorial by [Filippo Valsorda - THE ECB PENGUIN](https://blog.filippo.io/the-ecb-penguin/) to create the ecb party gophers.
For more details, please visit his post!

1. Exported the png to ppm, e.g., via GIMP or `convert 7TH_BIRTHDAY.png 7TH_BIRTHDAY.ppm`.
2. Extract the header of the image (shouldn't be encrypted. ;) ) ` head -n 3 7TH_BIRTHDAY.ppm > header.txt`
3. Extract the tail of the image which we will encrypt with ecb later on. `tail -n +4 7TH_BIRTHDAY.ppm > body.bin`
4. Encrypt with ECB `openssl enc -aes-128-ecb -nosalt -in body.bin -out body.ecb.bin`
5. Put the image together: `cat header.txt body.ecb.bin > 7TH_BIRTHDAY.ecb.ppm`
6. Exported the ppm to png, e.g., via GIMP or `convert 7TH_BIRTHDAY.ecb.ppm 7TH_BIRTHDAY.ecb2.png`. 


## Creating the CBC Image


Create a cbc encrypted version of the party gopher image:

````
$ openssl enc -aes-128-cbc -p --nosalt --pass pass:"Anna" -in body.bin -out body.cbc.bin
*** WARNING : deprecated key derivation used.
Using -iter or -pbkdf2 would be better.
key=BF4FCCD616251B678C56B9CB7A46819B
iv =1266853C180637642F5BC7D6B01F5554
$ cat header.txt body.cbc.bin > 7TH_BIRTHDAY.cbc.ppm
$ convert 7TH_BIRTHDAY.cbc.ppm 7TH_BIRTHDAY.cbc.png
````

Create another cbc encrypted version of the party gopher image 

````
$ openssl enc -aes-128-cbc -p --pass pass:"Anna" -in body.bin -out body.cbc2.bin -p --nosalt
   *** WARNING : deprecated key derivation used.
   Using -iter or -pbkdf2 would be better.
   key=63A3F2E1CD7BD1411221A75742F7C581
   iv =A309C1241AC5C24FCF4EA2D2E3964E79
$ cat header.txt body.cbc2.bin > 7TH_BIRTHDAY.cbc2.ppm
$ convert 7TH_BIRTHDAY.cbc2.ppm 7TH_BIRTHDAY.cbc2.png
````

Encrypts the image with the same IV:

````
$ openssl enc -aes-128-cbc -p --pass pass:"Anna" -in body.bin -out body.cbc3.bin -p --nosalt -iv A309C1241AC5C24FCF4EA2D2E3964E79 -K 63A3F2E1CD7BD1411221A75742F7C581
*** WARNING : deprecated key derivation used.
   Using -iter or -pbkdf2 would be better.
   key=63A3F2E1CD7BD1411221A75742F7C581
   iv =A309C1241AC5C24FCF4EA2D2E3964E79
$ cat header.txt body.cbc3.bin > 7TH_BIRTHDAY.cbc3.ppm
$ convert 7TH_BIRTHDAY.cbc3.ppm 7TH_BIRTHDAY.cbc3.png
````

Differs:
```
$ diff body.cbc2.bin body.cbc3.bin
$ diff body.cbc.bin body.cbc2.bin
Binary files body.cbc.bin and body.cbc2.bin differ
```

A bit more graphical:
````
$ compare 7TH_BIRTHDAY.cbc.png 7TH_BIRTHDAY.cbc2.png  -compose src -highlight-color seagreen 7TH_BIRTHDAY.cbc-vs-cbc2.png
$ compare 7TH_BIRTHDAY.cbc2.png 7TH_BIRTHDAY.cbc3.png  -compose src -highlight-color seagreen 7TH_BIRTHDAY.cbc2-vs-cbc3.png
````
