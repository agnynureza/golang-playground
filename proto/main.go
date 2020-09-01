package main

import (
	"bytes"
	"fmt"
	"os"
	"proto/model"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "hello world",
		Gender:   model.UserGender_FEMALE,
	}

	// var userList = &model.UserList{
	// 	List: []*model.User{
	// 		user1,
	// 	},
	// }

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	fmt.Println(garageListByUser)
	fmt.Println(user1)
	fmt.Println(user1.String())

	// ============================= as json string
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Println(jsonString)

	// json string to ====================== proto
	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Println(protoObject.String())

}
