package main

import (
	     "fmt"
	      consulapi "github.com/hashicorp/consul/api"
	    "log"
)

func main(){
      TestRegister()
      //TestDregister()
}

func TestRegister() {

    fmt.Println("test begin .")
    config := consulapi.DefaultConfig()
    config.Address = "109.254.2.139:8500"
   
    //config.Address = "localhost"
    fmt.Println("defautl config : ", config)
    client, err := consulapi.NewClient(config)
    if err != nil {
        log.Fatal("consul client error : ", err)
    }
    //创建一个新服务。
    registration := new(consulapi.AgentServiceRegistration)
    registration.ID = "redis_s2"
    registration.Name = "redis"
    registration.Port = 36379
    registration.Tags = []string{"redis"}
    registration.Address = "109.254.2.139"

    //增加check。
    check := new(consulapi.AgentServiceCheck)
    check.Script = fmt.Sprintf("redis-cli -h %s -p %d info | grep role:master || exit 2", registration.Address, registration.Port)
    //设置超时 5s。
    check.Timeout = "5s"
    //设置间隔 5s。
    check.Interval = "5s"
    //注册check服务。
    registration.Check = check
    log.Println("get check.Script:",check)

    err = client.Agent().ServiceRegister(registration)

    if err != nil {
        log.Fatal("register server error : ", err)
    }

}

func TestDregister(){


    fmt.Println("test begin .")
    config := consulapi.DefaultConfig()
    config.Address = "109.254.2.199:8500"
    //config.Address = "localhost"
    fmt.Println("defautl config : ", config)
    client, err := consulapi.NewClient(config)
    if err != nil {
        log.Fatal("consul client error : ", err)
    }
    err = client.Agent().ServiceDeregister("redis_s1")
    if err != nil {
        log.Fatal("register server error : ", err)
    }


}
