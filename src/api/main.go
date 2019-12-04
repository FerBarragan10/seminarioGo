package main
import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type Auto struct{
	velocidad int `json:"velocidad"`
	modelo int `json:"modelo"`
	marca string `json:"marca"`
	patente string `json:"patente"`
}
type agencia struct{
	autos[]Auto
}

func (a*Auto)acelerar(){
	a.velocidad+=10
} 

func newAuto(patent,marc string,velocid,model int)Auto{
	return Auto{
		patente:patent,
		modelo:model,
		velocidad:velocid,
		marca:marc,
	}
}

func GetHello(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "message": "hello world",
    })
}



func getAutosAgencia(agencia[]Auto){
	for i := 0; i < len(agencia); i++ {
		fmt.Println(agencia[i])
	}
}
func addAutoAgencia(agencia*[]Auto,Au Auto){
 *agencia=append(*agencia,Au)
}

func main() {
	var agencia []Auto	
	router := gin.Default()

	router.GET("/hello", GetHello)
	

if err := router.Run(); err != nil {
        fmt.Println(err)
    }

	

	//fmt.Println("hello world")
//	c :=otracarpeta.SumarValores(4 , 5)
	//au:=newAuto("sjf1234","volkswagen",200,1998)
	//fmt.Println("auto",au)
	
	//agencia = make([]Auto, 3)
	auto1:=newAuto("sjf1234","volkswagen",1998,200,)
	agencia=append(agencia,auto1)
	Au1:=newAuto("sjD234","fiat",1992,180,)
	addAutoAgencia(&agencia,Au1)
	//agencia=append(agencia,Au1)
	Au2:=newAuto("sdv294","reno",2002,210,)
	addAutoAgencia(&agencia,Au2)
	//agencia=append(agencia,Au2)
	Au3:=newAuto("ducunz","renault",2019,280,)
	addAutoAgencia(&agencia,Au3)
	getAutosAgencia(agencia)
	router.Run(":8080")

}
