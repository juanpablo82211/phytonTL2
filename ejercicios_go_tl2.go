package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("*************PYTHON SYNTAX**************\n")
	fmt.Println("ejercio 1:")
	fmt.Println("hello world <-- para poder imprimir en go necesitamos el paquete fmt el cual nos permite imprimir en consola texto ademas aqui solo se puede usar comillas dobles ")
	print("\nejercio 2:\n")
	if 5 > 2 {
		print("Five is greater than two!") //fmt.println acepta mas de 2 argumentos sin agregar una nueva linea en cambio print solo acepta 1 argumento a la vez
	}
	print("aqui para los if se necesitan los corchetes no como en python que son solo :\n")

	print("**************PYTHON COMMENTS**************\n") //comentarios
	print("\nejercio 1:")
	print("para comentar se necesitan // en vez de # como es en python")
	//this is a comment
	print("\nejercio 2:")
	print("para realizar un comentario de mas de una linea es necesario los simbolos /* y para cerrarlo */ en lugar de triple comillas como en python")
	/*this is a comment written in  more than just one line */

	print("\n**************PYTHON VARIABLES**************\n") //variables
	print("\nejercio 1:\n")
	var carname string = "Volvo"
	print("\naqui las variables se pueden declarar de 2 maneras var-nombre-tipo=  o nombre:= pero esta ultima solo dentro de funciones") //var nombreVariable tipoVariable" y también se puede utilizar la "declaración corta de variables" para declarar e inicializar variables en una sola línea.
	print(carname)
	print("\nejercicio 2:")
	var x int = 50
	print(x)
	print("\nejercicio 3:")
	var y int = 10
	print(x + y)
	print("\nejercicio 4:")
	x2 := 5
	y1 := 10
	z := x2 + y1
	print(z)
	print("\nejercicio 5:")
	print("2my-first_name = John no se debe hacer asi lo correcto seria myfirst_name= jhon")
	myfirst_name := "Jhon"
	print(myfirst_name)
	print("\nejercicio 6:")
	x3, y3, z3 := "orange", "orange", "orange"
	print("en go no es posible igual variables hay que definirlas por separado (en python se podria hacer x=y=z='orange')\n")
	print(x3 + " " + y3 + " " + z3 + " ")
	print("\nejercicio 7:")
	print("una funcion se declara de la sig manera: func+nombre()+tipo+{} no como en python que es  def+nombre+()\n")
	myfunc()

	print("\n**************PYTHON DATA TYPES**************\n") //tipos de datos
	print("ejercicio 1:\n")
	print("en go se necesita la libreria reflect para el comando reflect.TypeOf() para asi poder sacar el tipo de variable de la variable en cuestion no como en python que es type()\n")
	x10 := 5
	fmt.Println(reflect.TypeOf(x10))

	print("\nejercicio 2\n:")
	y10 := "hello world"
	fmt.Println(reflect.TypeOf(y10))
	print("\nejercicio 3\n:")
	x11 := 20.5
	fmt.Println(reflect.TypeOf(x11))

	print("\nejercicio 4\n:")
	print("en go para crear un slice de strings se usa []string , en cambio en python seria ['apple','banana','cherry']\n")
	x12 := []string{"apple", "banana", "cherry"}
	fmt.Println(reflect.TypeOf(x12))
	print("\nejercicio 5\n:")
	print("en go no se pueden crear tuplas para  crear una estructura que contenga multuiples valores se pueden usar slices,arrays o mapas\n")
	var array = [3]string{"apple", "banana", "cherry"}
	fmt.Println(reflect.TypeOf(array))
	print("\nejercicio 6:")
	print("\nen go no existe los dicts como en python,sino que existen los mapas los cuales  pueden ser cualquier tipo de datos, y se accede a ellos mediante las claves correspondientes.\n")
	informacion := make(map[string]string)
	informacion["nombre"] = "jhon"
	informacion["age"] = "36"

	fmt.Println(reflect.TypeOf(informacion))

	print("\nejercicio 7:")
	print("en go para declarar una variable como verdadera se debe hace en minusculas 'true'\n")
	x15 := true
	fmt.Println(reflect.TypeOf(x15))
	print("\n**************PYTHON NUMBERS **************\n")

	print("\nejercicio 1:\n")
	x20 := 5
	x20_float := float64(x20)                   //Estos valores son en realidad direcciones de memoria en hexadecimal que hacen referencia a las partes imaginarias del número complejo.
	print(x20_float, reflect.TypeOf(x20_float)) // y dos partes imaginarias: 0x8c6da0 y 0x8925c0

	print("\nejercicio 2:\n")
	x21 := 5.5
	x21_int := int(x21)
	print(x21_int, reflect.TypeOf(x21_int))
	print("\nejercicio 3:\n")
	x22 := 5
	x22_complex := complex(float64(x22), 0)
	print(x22_complex, reflect.TypeOf(x22_complex))

	print("\n**************PYTHON STRINGS **************\n")
	print("\nejercicio 1:")
	p := "hello world"
	print(len(p))
	print("\nejercicio 2:")
	txt := "hello world"
	p1 := txt[0] //lo da en codigo ascii
	print(p1)
	print("\nejercicio 3:")
	p2 := txt[2:5]
	print(p2)
	print("\nejercicio 4:")
	print("aqui el texto posee muchos espacios  con strings.trimspace() quita esos espacios\n")
	txt = "              hello world                          "

	print(txt)
	print("se convierte en: \n")
	p3 := strings.TrimSpace(txt) //en lugar de .strip() usar strings.TrimSpace()
	print(p3)
	print("\nejercicio 5:")
	txt = "hello world"
	p4 := strings.ToUpper(txt)
	print(p4)
	print("\nejercicio 6:")
	p5 := strings.ToLower(txt)
	print(p5)
	print("\nejercicio 7:")

	p6 := strings.Replace(txt, "h", "J", -1) // Si se establece en -1, se reemplazarán todas las ocurrencias. Si se establece en 0, no se realizará ningún reemplazo.
	print(p6)                                //Si se establece en un número positivo n, se reemplazarán las primeras n ocurrencias.
	print("\nejercicio 8:")
	age := 36
	txt = "My name is John, and I am %d years old" //Este código imprimirá la cadena "My name is John, and I am 36 years old". En la cadena de formato txt, %d es un marcador de posición para un valor entero (en este caso, la variable age)
	formatted := fmt.Sprintf(txt, age)             // Luego, la función fmt.Sprintf() inserta el valor de age en el marcador de posición y devuelve la nueva cadena de caracteres.
	fmt.Println(formatted)

	print("**************PYTHON BOOLEANS **************\n")

	print("\nejercicio 1:")
	print(10 > 9)
	print("--------- al ser 10>9 se imprimira un TRUE ya que esta afirmacion es verdadera")
	print("\nejercicio 2:")
	print(10 == 9)
	print("--------- al ser 10==9 se imprimira un FALSE ya que esta afirmacion es falsa")
	print("\nejercicio 3:")
	print(10 < 9)
	print("--------- al ser 10<9 se imprimira un FALSE ya que esta afirmacion es falsa")
	print("\nejercicio 4:\n")
	str := "true" //en go una cadena de caracterres no representa un bool, con este codigo podemos determinar si la cadena es verdadera o falsa "solo si es true o false de resto dara un error de syntaxis"
	b, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}
	print("\nejercicio 5:\n")
	k := 1 //en go debemos asignar el valor a una variable luego esa variable a otra variable para que pierda el tipo de variable actual y asi poder tomarla como un bool
	mbool := k == 1
	fmt.Println(bool(mbool)) // no como en python que se puede hacer directamente
	print("0<-false 1<- true")

	print("\n**************PYTHON OPERATORS **************\n")
	print("\nejercicio 1:")
	print(10 * 5)
	print("para multiplicar se usa el signo '*' en este caso se multiplico 10*5")
	print("\nejercicio 2:")
	print(10 / 2)
	print("para dividir se usa elñ signo '/' en este caso se dividio 10/2")
	print("\nejercicio 3:")
	fruits := []string{"apple", "banana"}
	if contains(fruits, "apple") {
		print("Yes, apple is a fruit!")
	}
	print("\n para verificar que apple esta en fruits debemos hacer una funcion de rango = al total de elementos o slice donde si el item en este caso manzana es igual a algun elemento nos devuelta true")
	print("\nejercicio 4:")
	if 5 != 10 {
		print("5 and 10 is not equal")
	}

	print("la expresion ! significa lo contrario de entonces si 5!=10 significa que 5 no es igual a 10")
	print("\nejercicio 5:")
	if 5 == 10 || 4 == 4 {
		print("At least one of the statements is true")
		print("go usa el algebra de bool, en este caso '||'significa al menos 1 de las n condiciones cumple ")
	}

}

var x string

func myfunc() int {
	x = "fantastic"
	fmt.Println(x)
	return 0
}

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
