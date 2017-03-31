# retool_test_app
A dummy repository for testing github.com/twitchtv/retool.

Has these tags:

 - tag 'v1', vranch `version1`: Has no dependencies. Program runs `fmt.Println("v1")`.
 - tag 'v2', branch `has_dep`: Depends on github.com/spenczar/retool_test_lib. Program runs `fmt.Println("v2")`.
 - tag 'v3', branch `remove_dep`: Removes dependency on retool_test_lib. Program runs `fmt.Println("v3")`.
 - tag 'v4', branch `master`: Program runs `fmt.Println("v4")`.
