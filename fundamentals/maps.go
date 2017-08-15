/*
Maps in Go
* Map keys have to have equality operator defined
  - Maps and slices do not have equality operator defined
* Maps are reference types
 - When a map is passed to a function, any changes inside the function
   modifies the original map as the reference is passed.
* Maps are not thread safe
 - Not recommended for concurrent programs
*/

/* Map Operations
** Insert
** Delete
** Update
** Check for existence
*/

package main

import "fmt"

func GetPrefix(name string) (prefix string) {
  // Creating a Map for names and their Prefixes
  // Initialization method 1:
  // Declare the types of the keys and values, in this case, key=> string
  // value => string
  var prefixMap map[string]string
  // Initialization using make - this is must as it will allocate the space to the map
  prefixMap = make(map[string]string)

  // Inserting into the map
  prefixMap["Bob"] = "Dr "
  prefixMap["Joe"] = "Mr "
  prefixMap["Mary"] = "Mrs "

  // Updating the value in the Map
  prefixMap["Bob"] = "Mr "

  // Check for existence
  // name, exists := prefixMap[name] ==> initializes two variables with name being
  // the key and a boolean 'exists' indicating whether the key exists or not.
  if name, exists := prefixMap[name]; exists {
    return prefixMap[name]
  } else {
    return "NA"
  }
}

func GetSurname(name string) (surname string) {
  // ShortHand Initialization of a map
  surnameMap := map[string]string {
    "Peter" : "Wills",
    "Doe" : "Jones",
    "Sam" : "Tarly", // Note: the trailing comma ',' is a must
  }
  // Deleting an entry in the Map
  delete(surnameMap, "Peter")

  return surnameMap[name]
}

func main()  {
  fmt.Println(GetSurname("John"))
  fmt.Println(GetPrefix("Hari"))
  fmt.Println(GetSurname("Peter"))

}
