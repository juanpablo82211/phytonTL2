package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type MyClass struct {
	x string
}
type MyClase struct {
	x int
}

type Persona struct {
	name string
	age  int
}

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

	fmt.Println("**************GO LISTS**************")

	fmt.Println("ejercicio 1:")
	frutas := []string{"apple", "banana", "cherry"} //toma el argumento en la posicion 1, la cual seria banana
	fmt.Println(frutas[1])

	fmt.Println("\nejercicio 2:")
	frutas[0] = "kiwi" //cambiamos la palabra en la posicion 0 por kiwi
	fmt.Println(frutas)

	fmt.Println("\nejercicio 3:")
	frutas = append(frutas, "orange") //agregamos con el parametro append la palabra orange
	fmt.Println(frutas)

	fmt.Println("\nejercicio 4:")
	frutas = append(frutas[:1], append([]string{"lemon"}, frutas[1:]...)...) //aqui lo que hacemos es agregar lemon a la posicion 1, primero copiamos la lista hasta la posicion1
	fmt.Println(frutas)                                                      //luego agregamos lemon y luego copiamos la lista desde la posicion 1 quedando entonces lemon incorporado

	fmt.Println("\nejercicio 5:")
	for i, fruit := range frutas {
		if fruit == "banana" {
			fruits = append(frutas[:i], frutas[i+1:]...) //borra banana de la lista
			break
		}
	}
	fmt.Println(frutas)

	fmt.Println("\nejercicio 6:")
	fmt.Println(frutas[len(frutas)-1]) // imprime el contenido de la lista marcado como -1 en este caso seria el ultimo de la lista (orange)

	fmt.Println("\nejercicio 7:")
	fmt.Println(frutas[2:5]) //imprime las palabras de la lista desde el 2 hasta el 5

	fmt.Println("\nejercicio 8:")
	fmt.Println(len(frutas)) //imprime la cantidad de strings que hay en la lista

	fmt.Println("**************GO TUPLES**************") //debemos tener en cuenta que tuplas en go son slices que se crean con la sintaxis []tipo_elemento{elementos}

	fmt.Println("\nejercicio 1:")
	frutas = []string{"apple", "banana", "cherry"} //imprime el elemento en la posicion 0
	fmt.Println(frutas[0])

	fmt.Println("\nejercicio 2:")
	frutas = []string{"apple", "banana", "cherry"} //imprime la cantidad de elementos
	fmt.Println(len(frutas))

	fmt.Println("\nejercicio 3:")
	frutas = []string{"apple", "banana", "cherry"} //imprime el primer elemento de atras hacia adelante es decir el ultimo elemento (cherry)
	fmt.Println(frutas[len(frutas)-1])

	fmt.Println("\nejercicio 4:")
	frutas = []string{"apple", "banana", "cherry", "orange", "kiwi", "melon", "mango"} //imprime los elementos de la posicion 2 a la 5
	fmt.Println(frutas[2:5])

	fmt.Println("**************GO SETS**************")

	fmt.Println("\nejercicio 1:")
	frutas1 := make(map[string]bool)
	frutas1["apple"] = true
	frutas1["banana"] = true
	frutas1["cherry"] = true

	if frutas1["apple"] {
		fmt.Println("Yes, apple is a fruit!")
	}

	fmt.Println("\nejercicio 2:")
	frutas1 = make(map[string]bool)
	frutas1["apple"] = true
	frutas1["banana"] = true
	frutas1["cherry"] = true //aqui añadimos orange no hay un .add para agregarla, debemos hacer el mapa y agregarla
	frutas1["orange"] = true

	fmt.Println("\nejercicio 3:")
	frutas1 = make(map[string]bool)
	frutas1["apple"] = true
	frutas1["banana"] = true
	frutas1["cherry"] = true
	moreFrutas := []string{"orange", "mango", "grapes"} //añadimos al mapa el slice more frutas con el ciclo fruta, el cual en el rango de moreFrutas añadara los elementos y los igualara a true
	for _, fruta := range moreFrutas {
		frutas1[fruta] = true //el guion bajo lo que hace es ignorar el indice,es decir, el numero entero que indica la posicion del elemento en el slice
	}
	for fruta, existe := range frutas1 {
		fmt.Printf("%s -> %v\n", fruta, existe) //ciclo para imprimir el mapa
	}

	fmt.Println("\nejercicio 4:")
	frutas1 = make(map[string]bool)
	frutas1["apple"] = true
	frutas1["banana"] = true
	frutas1["cherry"] = true
	delete(frutas1, "banana") //removemos banana con delete
	for fruta, existe := range frutas1 {
		fmt.Printf("%s -> %v\n", fruta, existe) //ciclo para imprimir el mapa
	}
	fmt.Println("\nejercicio 5:")
	frutas1 = make(map[string]bool)
	frutas1["apple"] = true
	frutas1["banana"] = true
	frutas1["cherry"] = true //discard es lo mismo que delete en go
	delete(frutas1, "banana")
	for fruta, existe := range frutas1 {
		fmt.Printf("%s -> %v\n", fruta, existe) //ciclo para imprimir el mapa
	}

	fmt.Println("**************GO MAPS**************\n") //en go no hay diccionarioos por lo que debemos hacer mapas para poder obtener los resultados

	fmt.Println("\nejercicio 1:")
	car := map[string]interface{}{ //interface proporcionan una manera flexible e trabajar con diferentes tipos de datos ene ste caso se trabaja con string y con int
		"brand": "Ford",
		"model": "Mustang", //mapa basico en go
		"year":  1964,
	}
	fmt.Println(car["model"]) //imprime el elemento guardado en model

	fmt.Println("\nejercicio 2:")
	car["year"] = 2020 //actualiza el año del mapa en este caso de 1964 a 2020
	fmt.Println(car["year"])

	fmt.Println("\nejercicio 3:")
	car["color"] = "red" //añade un nuevo elemento al mapa el cual es color el cual guarda red
	for carro, existe := range car {
		fmt.Printf("%s -> %v\n", carro, existe) //ciclo para imprimir el mapa
	}

	fmt.Println("el color del carro es:", car["color"])

	fmt.Println("\nejercicio 4:")
	delete(car, "model") //elimina del mapa car el elemento model
	fmt.Println(car)

	fmt.Println("\nejercicio 5:")
	for k := range car {
		delete(car, k) //elimina todos los elementos del mapa
	}
	fmt.Println(car)

	fmt.Println("**************GO IF...ELSE**************")

	fmt.Println("\nejercicio 1:")
	a := 50
	b1 := 10 //si a mayor que b imprima los valores de a y b y hello world
	if a > b1 {
		fmt.Println("a=", a, "b=", b1, "hello world")
	}

	fmt.Println("\nejercicio 2:")
	a = 50
	b1 = 10
	if a != b1 { //si a diferente de b imprime hello world
		fmt.Println("Hello World")
	}

	fmt.Println("\nejercicio 3:")
	a = 50
	b1 = 10 //si a igual que b1 imprime YES de lo contrario imprime NO
	if a == b1 {
		fmt.Println("Yes es igual")
	} else {
		fmt.Println("No, no es igual")
	}

	fmt.Println("\nejercicio 4:")
	a = 50
	b1 = 10 //declaramos las nuevas variables
	c := 5
	d := 10
	if a == b1 {
		fmt.Println("1") //si a igual a b1 imprime 1
	} else if a > b1 {
		fmt.Println("2") //si no es asi y a es mayor que b imprime 2
	} else {
		fmt.Println("3") //si es alguna otra opcion imprime 3
	}

	fmt.Println("\nejercicio 5:")
	if a == b1 && c == d { // si a es igual a b1 y c igual a d entonces imprime hello sino imprime bye bye
		fmt.Println("Hello")
	} else {

		fmt.Println("bye bye")
	} //recordemos que && necesita que ambas condiciones se cumplan en cambio || con que una condicion que se cumpla realiza la accion

	fmt.Println("\nejercicio 6:")
	if a == b1 || c == d { //si a es igual a b1 O c igual a d entonces imprime hello
		fmt.Println("Hello")
	}

	fmt.Println("\nejercicio 7:")
	if 5 > 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	print("**************GO WHILE (for) LOOPS**************\n") //aqui el while funciona con for

	fmt.Println("Ejercicio 1:")
	i := 1
	for i < 6 {
		fmt.Println(i) //imprime todos los numeros que sean menores que 6 empezando desde el 1
		i++
	}

	fmt.Println("\nEjercicio 2:")
	i = 1
	for i < 6 {
		if i == 3 { //si i llega al valor de 3 lo imprime y corta el ciclo
			fmt.Println(i)
			break
		}
		i++
	}

	fmt.Println("\nEjercicio 3:")
	i = 0
	for i < 6 {
		i++
		if i == 3 {
			continue //si i vale 3 se salta este valor y despues continua con el ciclo
		}
		fmt.Println(i)
	}

	fmt.Println("\nEjercicio 4:")
	i = 1
	for i < 6 {
		fmt.Println(i)
		i++
	}
	fmt.Println("i ya no es menor que 6")

	fmt.Println("**************GO FOR LOOPS**************")

	fmt.Println("\nejercicio 1:")
	fruits = []string{"apple", "banana", "cherry"}
	for _, x := range fruits { //imprime todos los elementos del slice con un ciclo
		fmt.Println(x)
	}

	fmt.Println("\nejercicio 2:")
	fruits2 := []string{"apple", "banana", "cherry"}
	for _, x := range fruits2 {
		if x == "banana" { // si x es igual a banana salta al siguiente ciclo
			continue
		}
		fmt.Println(x)
	}

	fmt.Println("\nejercicio 3:")
	for x := 0; x < 6; x++ {
		fmt.Println(x) //imprime los elementos mientras que x sea menor a 6
	}

	fmt.Println("\nejercicio 4:")
	fruits3 := []string{"apple", "banana", "cherry"}
	for _, x := range fruits3 {
		if x == "banana" { //aqui rompe el ciclo cuando se encuentra con la cadena de caracteres banana
			break
		}
		fmt.Println(x)
	}

	fmt.Println("**************GO FUNCTIONS**************")

	fmt.Println("\nejercicio 1:")
	myFunction()

	fmt.Println("\nejercicio 2:")
	myFunction2("juan", "perez")

	fmt.Println("\nejercicio 3:")
	result := myFunction3(5)
	fmt.Println(result)

	fmt.Println("\nejercicio 4:")
	myFunction4("juan", "jane", "david")

	fmt.Println("\nejercicio 5:")
	myFunction5(map[string]string{"lname": "smith"})

	print("**************GO LAMBDA**************\n")

	lambda := func(a int) int { //en go no existe lambda pero podemos hacer funciones anonimas el cual al ingresar 1 valor retorna ese mismo valor ingresa a retorna a
		return a
	}
	fmt.Println(lambda(5))

	print("**************PYTHON CLASSES**************\n") //en go no existen clases si no estructuras

	fmt.Println("\nejercicio 1:")

	clase := MyClass{x: "esto es una clase"}
	fmt.Println(clase.x)

	fmt.Println("\nejercicio 2:")
	clase2 := MyClase{x: 5}
	fmt.Println(clase2.x)

	fmt.Println("\nejercicio 3:")
	clase3 := Persona{name: "Juan", age: 30}
	clase4 := Persona{name: "Maria", age: 25}
	fmt.Println(clase3.name, clase3.age)
	fmt.Println(clase4.age)

	print("**************GO INHERITANCE**************\n") //go no tiene soporte para herencia de clases sin embargo podemos lograr el mismo resultado utilizando una composicion de estructuras

	type Persona struct {
		name string
		age  int
	}

	type Student struct {
		p  Persona
		id string
	}
	s1 := Student{Persona{"Juan", 20}, "12345"} //aqui stundent pide P pero p es persona por lo que nse dan los vales de la estructura persona en ester caso name y age
	fmt.Println(s1.p.name)                      //juan
	fmt.Println(s1.p.age)                       //20
	fmt.Println(s1.id)                          //12345

	print("**************GO MODULES**************\n") //para hacer modulos recordemos que realizamos un nuevo package en este caso "mymodule" aqui podemos guardar funciones, variables y demas
	print("\nejercicio 1:")
	print("import {\nmymodule}") //importamos un modulo al package main con import+ nombre de modulo entre comillas dobles
	print("\nejercicio 2:")
	print("import mymodule 'mx'") //importamos nuestro paquete y al lado ponemos el alias
	print("\nejercicio 3:")
	print("import mymodule")
	print("print(reflectTypeOf(mymodule))") //vemos que tiene adentro el paquete
	print("\nejercicio 4:")
	print("fmt.Println(mymodule.Person1.Name) \nfmt.Println(mymodule.Person1.Age)") //obtenemos el valor de person1 del  paquete mymodule
}

func myFunction() {
	fmt.Println("Hello from a function") //la funcion imprime hellow ferom a function cuando es llamada
}

func myFunction2(fname string, lname string) { // la funcion pide 2 valores un fname y un lname cuando son otorgados imprime el fname
	fmt.Println(fname)
}

func myFunction3(numero int) int {
	fmt.Println(numero) // la funcion pide un entero cuando es otorgado imprime el entero luego le suma 5 y retorna el entero+5
	return numero + 5
}

func myFunction4(kids ...string) {
	fmt.Println("The youngest child is", kids[2]) //la funcion pide una lista de strings y cuando se le otorga completa el mensaje con el elemento en la posicion 2
}

func myFunction5(kid map[string]string) { // la funcion pide un mapa y cuando es entregado completa el mensaje con lo que tiene el mapa en lname
	fmt.Println("His last name is", kid["lname"])
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
