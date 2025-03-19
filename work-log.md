# This file records raw thoughts, tasks, changes, and progress
## How I see the program functioning
- `rederb` with a path will just create a feed in the path (`rederb create`)
- `rederb` with nothing ~~should show help~~ should just create a feed in the current folder if there are any audio files in there
- `rederb create` should do the same as above: take the files in the current folder and create a feed according to defaults
- Empty, the site defaults to `http://localhost/audio-files-folder/feed-name.xml`
- Command line flags, that'll let me set site and path. (the final part of the path will always be the folder name)
    - for example: `http://localhost/fiction/audio-files-folder/feed-name.xml`
    - only one level for now. site -> category -> book-folder
- A config file, (yaml/env) that will hold defaults? or multiple choices?
    - go create --category=fiction will use fiction category in the url?
    - config file will also hold settings for feed, such as author, email, and site url





## Tasks as I see them
- [X] Learn how pull requests work in go land
- [X] Find dependencies, like I use in Python  
  viper(MIT)=python-dotenv  
  dhowden/tag(BSD-2)=tinytag  
  gorilla/feeds(BSD-3)=feedgen  
  *and brand new:* goreleaser for releases 
  ~~Nothing set in stone, yet~~  *(Update, 2025-03-17: Now it is)*
- [X] Update Project LICENSE once I figure out above (BSD 2 Clause)
- [X] Learn the basic go toolchain. does not have to be perfect. Learn industry standards later. Right now I desperately need derb to work on the Pi
- [X] Figure out how to read in environment variables from a .env file
- [X] Figure out to read metadata from audio files
- [X] Change the whole project into a Cobra app (2025-02-23) (done: 2025-02-24)
- [X] Create an app that just saves config (2025-03-15)
- [X] Get path from command line
- [X] verify path,
- [X] get stem from path and tag to feed url
- [X] split path to get end folder to tag on to url
- [X] and prints things out
- [X] figure out how to read in the contents of a directory and print filenames to screen
- [X] figure out how to import a tag package and get it to read the filenames you give it and print out the metadata it gets
- [X] Figure out how to write images from metadata to disk
- [X] Write slowly. Learn how types work as you go.
- [X] Learn how to create a feed
- [X] Accept any folder, read the files in there and then write the feed back in there
- [X] check only for audio files and print those
- [X] print them sorted? so i can pass them along in that order to the rss library
- [X] do we need to sort by filename? or track num? figure that out. *(Narrator: we used track numbers)*
- [X] Create a feed
- [X] First run!
- [X] Learn how github releases work, so I can distribute stuff
- [X] if there is a cover, leave it alone
- [ ] better changelog in goreleaser
- [ ] Add sample config file and include instructions on placement
- [ ] Integrate feedback from friends and family
- [ ] a `rederb init` command that creates a new config file in the current folder? or saves to `~/.config/rederb/rederb.yaml`
- [ ] Tests

---