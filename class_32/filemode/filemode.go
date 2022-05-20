package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))
	fmt.Println(os.FileMode(0707))
}

/*Each file has three sets of permissions, affecting three different classes of users.
The first set of permissions applies only to the user that owns the file. (By
default, your user account is the owner of any files you create.) The second set
of permissions is for the group of users that the file is assigned to. And the third
set applies to other users on the system that are neither the file owner nor part of
the file’s assigned group.
FileMode*/