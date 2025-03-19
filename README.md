# ReDerb

***Unfinished WIP! DO NOT USE***  
[Derb](https://github.com/jasonbraganza/derb) is Python. ReDerb is Go.

## What is ReDerb
ReDerb is a small, niche utility that takes a path to a directory full of audio files
and creates a RSS feed of the audio files in there to serve as a podcast feed.  
I need to serve audio files to my family, and serving them up as a feed, 
so that they could subscribe to it in their podcast players seemed like a good idea.

## Requirements
- Works only on 64-bit Linux
- A folder full of audio files that have audio tags in them.  
 (if not, use something like Ex Falso, Puddletag, and Kid3 on Linux or Tag & Rename on Windows (paid software, I’m not aware of opensource options on Windows) to tag them the way you want.)

### With the following constraints …
- needs feed settings, that you will provide in ReDerb’s settings file
- with only a single directory. ReDerb does not recurse into subdirectories

## Why ReDerb
I want to learn Go.  
So rewriting something I use frequently is good motivation  
The code, to my mind, is both not idiomatic (because of ignorance, new programmer) as well as wildly overkill. (I have a new hammer / throw everything at the wall to see what sticks [^1] )

[^1]: I am the undisputed champ of mixed metaphors though.

## Tasks as I see them
- now in the [work log](work-log.md)

## Contributing
This is a really niche little program for personal use and I wish to keep it that way.  
Also, I wrote this primarily to help me learn Go.  
So rather than actual code contributions, I’d love for you to review my code and tell me
- what you would see improved in terms of the code written
- while also telling / teaching me *why* you think it should be changed the way you suggest it should be  
  
In a nutshell, I’d love your contribution, to be teaching me.

## License
[BSD-2-Clause license.](https://opensource.org/license/bsd-2-clause)
See [LICENSE](LICENSE)

## Gratitude
- [David Howden](https://github.com/dhowden) for [tag](https://github.com/dhowden/tag)
- [Steve Francia](https://spf13.com/), [Márk Sági-Kazár](https://github.com/sagikazarmark) & Co. for [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)
- [Jason Moiron](https://github.com/jmoiron) and all the other friendly gorillas behind [Gorilla Feeds](https://github.com/gorilla/feeds)