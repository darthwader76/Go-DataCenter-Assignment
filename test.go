
//Parth Rajendra Doshi Early Warning Intern - DevOps Engineer Exercise


//This is the Golang program where we use predefined input from the user and assign the VMs dynamically to the predefined clusters
//based on the following criteria:
//Cluster - should be picked up based on environment and OS type
//Hosts   - should be picked up randomly. 
//Datastore - should be picked up based on most available free space on a datastore. 
//          - skip the datastore which has maintenancemode flag is set to 'yes'
//Networks - should be picked up based on IP address of the VM. compare IP address of the VM with VLAN address. 
//           Please ignore the 4th octet of IP address for comparison.


//I have used 3 data types namely Datastore, Cluster and Virtual_Machine using struct, I have defined maps, arrays as needed for initialization.
//I have used control loops namely if else, for for decision making and iteration purposes. My code currently won't return any exception coz it has predefined input, but in case we change the input, it could throw exception via function calls for type exceptions.
//I have used recover, defer and panic for exception handling in remove_index to show an example.
//I have used predefined utility functions to remove an randomly selected host, select datastore with max availability and print function to print cluster before assignment and after assignment using defer.

//The code could be optimised in terms of restructuring the repetitive blocks for cluster assignment. I wanted to finish the assignment as soon as possible and submit it and hence I have not done it right now.
//Further, input can be given through files.

//Instructions for executing this code
//  - change your directory to the directory of this file
//  - execute "go run test.go" in your cmd
//Expected Output is present in the output.txt file and at the end of this file too in comments.

package main

import (
    "fmt"
    "strings"
    "math/rand"
    "reflect"
)

type Datastore struct{
    name string
    capacity uint
    maintenancemode bool
}

type Cluster struct{
    name string
    ostype string
    environment string
    host[] string
    network map[string]string
    datastore map[int]Datastore
}

type Virtual_Machine struct{
    name string
    ostype string
    environment string
}

func main(){
    var cluster1 Cluster
    var cluster2 Cluster
    cluster1.name = "cluster-dcw-lnx-dev"
    cluster1.ostype = "linux"
    cluster1.environment = "dev"
    cluster1.host = []string {
        "host-dcw-lnx-dev-1",
        "host-dcw-lnx-dev-2",
        "host-dcw-lnx-dev-3",
        "host-dcw-lnx-dev-4",
        "host-dcw-lnx-dev-5",
    }
    cluster1.network = map[string]string{
        "10.16.153":"153_dev_pg_1",
        "10.16.154":"154_dev_pg_2",
        "10.16.179":"179_dev_pg_3",
        "10.16.180":"180_dev_pg_4",
    }
    cluster1.datastore = map[int]Datastore{
        3820875461632:Datastore{"ds-lnx-dev-4",4398046511104,false},
        3610875461632:Datastore{"ds-lnx-dev-3",4398046511104,true},
        2794586464256:Datastore{"ds-lnx-dev-2",5717460467712,true},
        1854907772928:Datastore{"ds-lnx-dev-1",8796093022208,true},
    }
    cluster2.name = "cluster-dcw-win-qa"
    cluster2.ostype = "windows"
    cluster2.environment = "qa"
    cluster2.host =[]string {
        "host-dcw-win-qa-1",
        "host-dcw-win-qa-2",
        "host-dcw-win-qa-3",
        "host-dcw-win-qa-4",
        "host-dcw-win-qa-5",
    }
    cluster2.network = map[string]string{
        "10.100.8":"8_qa_pg_1",
        "10.100.52":"52_qa_pg_2",
        "10.100.53":"53_qa_pg_3",
        "10.100.57":"57_qa_pg_4",
    }
    cluster2.datastore = map[int]Datastore{
        3441869148160:Datastore{"ds-win-qa-1",7311752327168,true},
        3651155738624:Datastore{"ds-win-qa-2",9400824418304,true},
        3610875461632:Datastore{"ds-win-qa-3",4398046511104,false},
        4044135927808:Datastore{"ds-win-qa-4",5277655814144,true},
    }
    fmt.Println("----------------------------------------------------------------------------------")
    printCluster(cluster1)
    fmt.Println("----------------------------------------------------------------------------------")
    printCluster(cluster2)
    fmt.Println("----------------------------------------------------------------------------------")
    virtmachines := map[string]Virtual_Machine{
        "10.16.153.12":Virtual_Machine{"vm5533dlv","linux","dev"},
        "10.16.179.4":Virtual_Machine{"vm7783dlv","linux","dev"},
        "10.16.180.7":Virtual_Machine{"vm2233dlv","linux","qa"},
        "10.100.8.9":Virtual_Machine{"vm2355dwv","windows","qa"},
        "10.100.57.2":Virtual_Machine{"vm2511dwv","windows","qa"},
        "10.100.60.6":Virtual_Machine{"vm2519dwv","windows","qa"},
    }
    fmt.Println()
    fmt.Println("----------------------------------------------------------------------------------")
    fmt.Println()
    //fmt.Printf( "Cluster1 Hosts : %s\n", strings.Join(cluster1.host,", "));
    //fmt.Printf( "Cluster2 Hosts : %s\n", strings.Join(cluster2.host,", "));
    defer func() {
        //fmt.Printf( "Cluster1 Hosts : %s\n", strings.Join(cluster1.host,", "));
        //fmt.Printf( "Cluster2 Hosts : %s\n", strings.Join(cluster2.host,", "));
        fmt.Println("----------------------------------------------------------------------------------")
        printCluster(cluster1)
        fmt.Println("----------------------------------------------------------------------------------")
        printCluster(cluster2)
        fmt.Println("----------------------------------------------------------------------------------")
    
    }()
    for vm := range virtmachines {
        flag := false
        flag_datastore := false
        if virtmachines[vm].ostype == cluster1.ostype && virtmachines[vm].environment == cluster1.environment {
            
            //allocate(cluster1, virtmachines,vm, flag)
            fmt.Println("Name:",virtmachines[vm].name)
            fmt.Println("OStype:",virtmachines[vm].ostype)
            fmt.Println("IP:",vm)
            fmt.Println("Environment:",virtmachines[vm].environment)
            fmt.Println("Cluster:",cluster1.name)
            if len(cluster1.host)==0{
                fmt.Println("Host: No host available for given",virtmachines[vm].name,"in the",cluster1.ostype,"environment")
            }else{
                i := rand.Intn(len(cluster1.host))
                fmt.Println("Host:",cluster1.host[i])
                defer func() {
                    if r:=recover(); r!=nil{
                        fmt.Println("Recovered in cluster1", r)
                    }
                }()
                cluster1.host = RemoveIndex(cluster1.host,i)

            }
            //fmt.Println(max(cluster.datastore))
            if len(cluster1.datastore)==0{
                fmt.Println("Datastore: No datastore available for given",virtmachines[vm].name,"in the",cluster1.ostype,"environment")
            }else{
                key := max(cluster1.datastore)
                for cluster1.datastore[key].maintenancemode != true{
                    delete(cluster1.datastore, key)
                    key = max(cluster1.datastore)
                    if len(cluster1.datastore) == 0{
                        break
                        flag_datastore = true
                    }
                }
                if !flag_datastore{
                    fmt.Println("Datastore:",cluster1.datastore[key].name)
                    delete(cluster1.datastore, key)    
                }else{
                    fmt.Println("Datastore: No datastore available for given",virtmachines[vm].name,"in the",cluster1.ostype,"environment")
                }
                
                
            }
            for k  := range cluster1.network{
                if strings.Contains(vm, k){
                    fmt.Println("Network:",cluster1.network[k])
                    flag = true
                    break
                }

            }
            if !flag {
                fmt.Println("Network: No VLAN network found for given IP",vm,"for",virtmachines[vm].name)
            }


        } else if virtmachines[vm].ostype == cluster2.ostype && virtmachines[vm].environment == cluster2.environment {
            //allocate(cluster2, virtmachines,vm,flag)
            fmt.Println("Name:",virtmachines[vm].name)
            fmt.Println("OStype:",virtmachines[vm].ostype)
            fmt.Println("IP:",vm)
            fmt.Println("Environment:",virtmachines[vm].environment)
            if len(cluster2.host)==0{
                fmt.Println("Host: No host available for given",virtmachines[vm].name,"in the",cluster2.ostype,"environment")
            }else{
                i := rand.Intn(len(cluster2.host))
                fmt.Println("Host:",cluster2.host[i])
                defer func() {
                    if r:=recover(); r!=nil{
                        fmt.Println("Recovered in cluster2", r)
                    }
                }()
                cluster2.host = RemoveIndex(cluster2.host,i)

            }
            //fmt.Println(max(cluster.datastore))
            if len(cluster2.datastore)==0{
                fmt.Println("Datastore: No datastore available for given",virtmachines[vm].name,"in the",cluster2.ostype,"environment")
            }else{
                key := max(cluster2.datastore)
                for cluster2.datastore[key].maintenancemode != true{
                    delete(cluster2.datastore, key)
                    key = max(cluster2.datastore)
                    if len(cluster2.datastore) == 0{
                        break
                        flag_datastore = true
                    }
                }
                if !flag_datastore{
                    fmt.Println("Datastore:",cluster2.datastore[key].name)
                    delete(cluster2.datastore, key)    
                }else{
                    fmt.Println("Datastore: No datastore available for given",virtmachines[vm].name,"in the",cluster2.ostype,"environment")
                }
                
                
            }

            for k  := range cluster2.network{
                if strings.Contains(vm, k){
                    fmt.Println("Network:",cluster2.network[k])
                    flag = true
                    break
                }

            }
            if !flag {
                fmt.Println("Network: No VLAN network found for given IP",vm,"for",virtmachines[vm].name)
            }


        } else {
            fmt.Printf("No %s cluster found for %s in %s environment.\n",virtmachines[vm].ostype,virtmachines[vm].name,virtmachines[vm].environment)
        }
        fmt.Println()
        fmt.Println("----------------------------------------------------------------------------------")
        fmt.Println()
    }

}

//utility function used to eliminate randomly selected hosts so that they are unavailable for selection in the future.
func RemoveIndex(s []string, index int) []string {
    if reflect.TypeOf(index).Kind()!=reflect.Int{
        panic(fmt.Sprintf("%v",index))
    }
    return append(s[:index], s[index+1:]...)
}

//Constants initialized for max function
const (
    maxInt = int(^uint(0) >> 1)
    minInt = -maxInt - 1
)

//function to return the max key from DataStore
func max(numbers map[int]Datastore) int {
    var maxNumber int
    for maxNumber = range numbers {
        break
    }
    for n := range numbers {
        if n > maxNumber {
            maxNumber = n
        }
    }
    return maxNumber
}

//Dummy function tried but didn't work. Need to understand scope of variables properly.
// func allocate(cluster Cluster, virtmachines map[string]Virtual_Machine, vm string, flag bool){


//     fmt.Println(virtmachines[vm].name)
//     fmt.Println(virtmachines[vm].ostype)
//     fmt.Println(vm)
//     fmt.Println(virtmachines[vm].environment)
//     i := rand.Intn(len(cluster.host))
//     fmt.Println(cluster.name)
//     fmt.Println(cluster.host[i])
//     cluster.host = RemoveIndex(cluster.host,i)

//     //fmt.Println(max(cluster.datastore))
//     key := max(cluster.datastore)
//     for cluster.datastore[key].maintenancemode != true{
//         delete(cluster.datastore, key)
//         key = max(cluster.datastore)    
//     }
//     fmt.Println(cluster.datastore[key].name)
//     delete(cluster.datastore, key)
//     for k  := range cluster.network{
//         if strings.Contains(vm, k){
//             fmt.Println(cluster.network[k])
//             flag = true
//             break
//         }

//     }
//     if !flag {
//         fmt.Println("No VLAN network found for given IP",vm,"for",virtmachines[vm].name)
//     }


// }


func printCluster( cluster Cluster ) {
    fmt.Printf( "Cluster name : %s\n", cluster.name);
    fmt.Printf( "Cluster ostype : %s\n", cluster.ostype);
    fmt.Printf( "Cluster environment : %s\n", cluster.environment);
    fmt.Println( "Cluster network :",cluster.network);
    fmt.Println( "Cluster datastore :",cluster.datastore);
    fmt.Printf( "Cluster Hosts : %s\n", strings.Join(cluster.host,", "));
}


// Runtime Error Handling
// Exception Handling

//Appreciate your time and consideration in understanding the code. It is my first attempt in Go Language and I hope I was able to perform as expected.
//Apologises for any unintentional mistakes.

//Expected Output
/*
C:\Users\ParthD\Documents\Early Warning Assignment>go run test.go
----------------------------------------------------------------------------------
Cluster name : cluster-dcw-lnx-dev
Cluster ostype : linux
Cluster environment : dev
Cluster network : map[10.16.153:153_dev_pg_1 10.16.154:154_dev_pg_2 10.16.179:179_dev_pg_3 10.16.180:180_dev_pg_4]
Cluster datastore : map[1854907772928:{ds-lnx-dev-1 8796093022208 true} 2794586464256:{ds-lnx-dev-2 5717460467712 true} 3610875461632:{ds-lnx-dev-3 4398046511104 true} 3820875461632:{ds-lnx-dev-4 4398046511104 false}]
Cluster Hosts : host-dcw-lnx-dev-1, host-dcw-lnx-dev-2, host-dcw-lnx-dev-3, host-dcw-lnx-dev-4, host-dcw-lnx-dev-5
----------------------------------------------------------------------------------
Cluster name : cluster-dcw-win-qa
Cluster ostype : windows
Cluster environment : qa
Cluster network : map[10.100.52:52_qa_pg_2 10.100.53:53_qa_pg_3 10.100.57:57_qa_pg_4 10.100.8:8_qa_pg_1]
Cluster datastore : map[3441869148160:{ds-win-qa-1 7311752327168 true} 3610875461632:{ds-win-qa-3 4398046511104 false} 3651155738624:{ds-win-qa-2 9400824418304 true} 4044135927808:{ds-win-qa-4 5277655814144 true}]
Cluster Hosts : host-dcw-win-qa-1, host-dcw-win-qa-2, host-dcw-win-qa-3, host-dcw-win-qa-4, host-dcw-win-qa-5
----------------------------------------------------------------------------------

----------------------------------------------------------------------------------

Name: vm2511dwv
OStype: windows
IP: 10.100.57.2
Environment: qa
Host: host-dcw-win-qa-2
Datastore: ds-win-qa-4
Network: 57_qa_pg_4

----------------------------------------------------------------------------------

Name: vm2519dwv
OStype: windows
IP: 10.100.60.6
Environment: qa
Host: host-dcw-win-qa-5
Datastore: ds-win-qa-2
Network: No VLAN network found for given IP 10.100.60.6 for vm2519dwv

----------------------------------------------------------------------------------

Name: vm5533dlv
OStype: linux
IP: 10.16.153.12
Environment: dev
Cluster: cluster-dcw-lnx-dev
Host: host-dcw-lnx-dev-3
Datastore: ds-lnx-dev-3
Network: 153_dev_pg_1

----------------------------------------------------------------------------------

Name: vm7783dlv
OStype: linux
IP: 10.16.179.4
Environment: dev
Cluster: cluster-dcw-lnx-dev
Host: host-dcw-lnx-dev-5
Datastore: ds-lnx-dev-2
Network: 179_dev_pg_3

----------------------------------------------------------------------------------

No linux cluster found for vm2233dlv in qa environment.

----------------------------------------------------------------------------------

Name: vm2355dwv
OStype: windows
IP: 10.100.8.9
Environment: qa
Host: host-dcw-win-qa-3
Datastore: ds-win-qa-1
Network: 8_qa_pg_1

----------------------------------------------------------------------------------

----------------------------------------------------------------------------------
Cluster name : cluster-dcw-lnx-dev
Cluster ostype : linux
Cluster environment : dev
Cluster network : map[10.16.153:153_dev_pg_1 10.16.154:154_dev_pg_2 10.16.179:179_dev_pg_3 10.16.180:180_dev_pg_4]
Cluster datastore : map[1854907772928:{ds-lnx-dev-1 8796093022208 true}]
Cluster Hosts : host-dcw-lnx-dev-1, host-dcw-lnx-dev-2, host-dcw-lnx-dev-4
----------------------------------------------------------------------------------
Cluster name : cluster-dcw-win-qa
Cluster ostype : windows
Cluster environment : qa
Cluster network : map[10.100.52:52_qa_pg_2 10.100.53:53_qa_pg_3 10.100.57:57_qa_pg_4 10.100.8:8_qa_pg_1]
Cluster datastore : map[]
Cluster Hosts : host-dcw-win-qa-1, host-dcw-win-qa-4
----------------------------------------------------------------------------------
*/