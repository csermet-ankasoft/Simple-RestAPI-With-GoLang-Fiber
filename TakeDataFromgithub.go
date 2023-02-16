package main

import "log"

// Takes data from github and returns the value as string
func TakeDataFromgithub(URL string, FileName string) string {
	log.Printf("%v Started to fetch data from Github", APINAME)

	fs := memfs.New()
	//Authenticate and clone the repository
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: URL,
	})
	if err != nil {
		log.Printf("%v Repo could not find. Error: %v", APINAME, err.Error())
	} else {
		log.Printf("%v Fetched Repository :  %v", APINAME, repo)
	}
	//Reads the repository as byte
	file, err := fs.Open(FileName)
	if err != nil {
		log.Printf("%v File could not find. Error: %v", APINAME, err.Error())
	} else {
		log.Printf("%v Fetched Repositories File :  %v", APINAME, file.Name())
	}

	defer func(file billy.File) {
		err := file.Close()
		if err != nil {
			log.Printf("%v Cannot close file. Error: %v", APINAME, err.Error())
		}
	}(file)
	//turns into a string
	return ParseData(file)
}
