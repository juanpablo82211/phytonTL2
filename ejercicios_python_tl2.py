print('*************PYTHON SYNTAX**************\n')  #syntaxis
print('ejercio 1:')
print('hello world')
print('ejercio 2:')
if 5 > 2:
 print("Five is greater than two!")
print('**************PYTHON COMMENTS**************\n')   #comentarios
print('\nejercio 1:')
print('aparece un comentario')
#this is a comment
print('\nejercio 2:')
print('aparece un comentario en mas de una linea')
'''this is a comment written in  more than just one line  '''
print('**************PYTHON VARIABLES**************\n')  #variables
print('\nejercio 1:')
carname="Volvo"
print(carname)
print("\nejercicio 2:")
x=50
print(x)
print("\nejercicio 3:")
x1= 50
y=10
print(x+y)
print("\nejercicio 4:")
x2=5
y1=10
z=x+y
print(z)
print("\nejercicio 5:")
print("2my-first_name = John no se debe hacer asi lo correcto seria myfirst_name= jhon")
myfirst_name="Jhon"
print(myfirst_name)
print("\nejercicio 6:")
x3=y3=z3="orange"
print(x3 + y3 + z3 )
print("\nejercicio 7:")

def myfunc():
 global x
 x= "fantastic"
 print(x)

 myfunc()

 print('**************PYTHON DATA TYPES**************\n')  #tipos de datos
print("ejercicio 1:")
x10=5
print(type(x10))
print("\nejercicio 2:")
y10="hello world"
print(type(y10))
print("\nejercicio 3:")
x11=20.5
print(type(x11)
)
print("\nejercicio 4:")
x12=["apple","banana","cherry"]
print(type(x12))
print("\nejercicio 5:")
x13 = ("apple", "banana", "cherry")
print(type(x13))
print("\nejercicio 6:")
x14 = {"name" : "John", "age" : 36} 
print(type(x14))
print("\nejercicio 7:")
x15=True
print(type(x15))

print("**************PYTHON NUMBERS **************\n")

print("\nejercicio 1:")
x20=5
x20=float(x20)
print(x20,type(x20))
print("\nejercicio 2:")
x21=5.5
x21=int(x21)
print(x21,type(x21))
print("\nejercicio 3:")
x22=5
x=complex(x22)
print(x22,type(x22))

print("**************PYTHON STRINGS **************\n")
print("\nejercicio 1:")
p="hello world"
print(len(p)) 
print("\nejercicio 2:")
txt="hello world"
p1=txt[0]
print(p1)
print("\nejercicio 3:")
p2=txt[2:5]
print(p2)
print("\nejercicio 4:")
print(  "aqui el texto posee muchos espacios  con .strip() quita esos espacios\n")
txt="              hello world                          "

print(txt)
print("se convierte en: \n")
p3= txt.strip()
print(p3)
print("\nejercicio 5:")
txt=txt="hello world"
p4=txt.upper()
print(p4)
print("\nejercicio 6:")
p5=txt.lower()
print(p5)
print("\nejercicio 7:")

p6=txt=txt.replace("h","J")
print(p6)
print("\nejercicio 8:")
age = 36
txt = "My name is John, and I am {}"
print(txt.format(age))

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
print("\nejercicio 4:")
print(bool("abc")) 
print("cualquier string es True a menos que este vacio ")
print("\nejercicio 5:")
print(bool(1))
print("0<-false 1<- true")

print("**************PYTHON OPERATORS **************\n")
print("\nejercicio 1:")
print(10*5)
print("para multiplicar se usa el signo '*' en este caso se multiplico 10*5")
print("\nejercicio 2:")
print(10/2)
print("para dividir se usa elñ signo '/' en este caso se dividio 10/2")
print("\nejercicio 3:")
fruits = ["apple", "banana"]
if "apple" in fruits:
  print("Yes, apple is a fruit!")
  print("aqui si el objeto apple esta en el conjunto fruits me lo confirma, se usa con la expresion 'in'")
print("\nejercicio 4:")
if 5 !=10:
  print("5 and 10 is not equal")
print("la expresion ! significa lo contrario de entonces si 5!=10 significa que 5 no es igual a 10")
print("\nejercicio 5:")
if 5 == 10 or 4 == 4:
  print("At least one of the statements is true")
  print("python usa el algebra de bool, en este caso 'or'significa al menos 1 de las n condiciones cumple ")

print("**************PYTHON LISTS **************\n")
print("\nejercicio 1:")
fruits = ["apple", "banana", "cherry"]
print(
fruits[1]
)
print("\nejercicio 2:")
fruits = ["apple", "banana", "cherry"]
fruits[0] = "kiwi"
print("\nejercicio 3:")
fruits = ["apple", "banana", "cherry"]
fruits.append("orange")
print("\nejercicio 4:")
fruits = ["apple", "banana", "cherry"]
fruits.insert(1,"lemon")
print("\nejercicio 5:")
fruits = ["apple", "banana", "cherry"]
fruits.remove("banana")
print("\nejercicio 6:")
fruits = ["apple", "banana", "cherry"]
print(fruits[-1])
print("\nejercicio 7:")
fruits = ["apple", "banana", "cherry", "orange", "kiwi", "melon", "mango"]
print(fruits[2:5])
print("\nejercicio 8:")

fruits = ["apple", "banana", "cherry"]
print(len(fruits))

print("**************PYTHON TUPLES**************\n")

print("\nejercicio 1:")
fruits = ("apple", "banana", "cherry")
print(fruits[0])
print("\nejercicio 2:")
fruits = ("apple", "banana", "cherry")
print(len(fruits))
print("\nejercicio 3:")
fruits = ("apple", "banana", "cherry")
print(fruits[-1])
print("\nejercicio 4:")
fruits = ("apple", "banana", "cherry", "orange", "kiwi", "melon", "mango")
print(fruits[2:5])

print("**************PYTHON SETS**************\n")

print("\nejercicio 1:")
fruits = {"apple", "banana", "cherry"}
if "apple" in fruits: print("Yes, apple is a fruit!")

print("\nejercicio 2:")
fruits = {"apple", "banana", "cherry"}
fruits.add("orange")

print("\nejercicio 3:")
fruits = {"apple", "banana", "cherry"}
more_fruits = ["orange", "mango", "grapes"]
fruits.update(more_fruits)

print("\nejercicio 4:")
fruits = {"apple", "banana", "cherry"}
fruits.remove("banana")

print("\nejercicio 5:")
fruits = {"apple", "banana", "cherry"}
fruits.discard("banana")

print("**************PYTHON DiCTIONARIES**************\n")

print("\nejercicio 1:")
car =	{
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
print(car.get("model"))
print("\nejercicio 2:")
car =	{
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
car["year"] = 2020
print(car.get("year"))
print("\nejercicio 3:")
car =	{
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
car["color"]="red"
print(car.get("color"))

print("\nejercicio 4:")
car =	{
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
car.pop("model")
print(car)

print("\nejercicio 5:")
car =	{
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
car.clear()
print(car)

print("**************PYTHON IF...ELSE**************\n")

print("\nejercicio 1:")
a = 50
b = 10
if a > b: print("a=",a,"b=",b,"hello world")

print("\nejercicio 2:")
a = 50
b = 10
if a !=b: print("Hello World")
print("\nejercicio 3:")
a = 50
b = 10
if a==b:print("Yes") 
else: print("No")
print("\nejercicio 4:")
a = 50
b = 10
c = 5
d = 10

if a==b:print("1")
elif a > b :print("2")
else: print("3")
print("\nejercicio 5:")
if a == b and c == d:
  print("Hello")
print("\nejercicio 6:")
if a == b or c == d:
 print("Hello")
print("\nejercicio 7:")
print("Yes") if 5 > 2 else print("No")

print("**************PYTHON WHILE LOOPS**************\n")

print("\nejercicio 1:")
i = 1
while i < 6:
 print(i) 
 i += 1
print("\nejercicio 2:")
i = 1
while i < 6:
    if i == 3:
        print(i)
        break
    i += 1


print("\nejercicio 3:")
i= 0
while i< 6:
  i+=1
  if i==3:
   continue
  print(i)
print("\nejercicio 4:")

i=1
while i< 6:
  print(i)
  i+=1
else:
  print ("i is no longer than 6") 

print("**************PYTHON WHILE LOOPS**************\n")

print("\nejercicio 1:")
fruits = ["apple", "banana", "cherry"]
for x in fruits:
   print(x)
print("\nejercicio 2:")
fruits = ["apple", "banana", "cherry"]
for x in fruits:
  if x == "banana":
    continue
  print(x)
print("\nejercicio 3:")
for x in range(6):
 print(x)
print("\nejercicio 4:")

fruits = ["apple", "banana", "cherry"]
for x in fruits:
  if x == "banana":
    break
  print(x)

  print("**************PYTHON FUNCTIONS**************\n")

print("\nejercicio 1:")
def my_function():
  print("Hello from a function")
my_function()
print("\nejercicio 2:")
def my_function(fname, lname):
  print(fname)
my_function("juan","perez")
print("\nejercicio 3:")
def my_function(numero):
 print(numero)
 return numero+5
my_function(5)

print("\nejercicio 4:")
def my_function(*kids): 
  print("The youngest child is " + kids[2])
my_function("juan","jane","david")
print("\nejercicio 5:")
def my_function(**kid): 
  print("His last name is " + kid["lname"])
my_function(lname="smith")

  
print("**************PYTHON LAMBDA**************\n")

x = lambda a:a
print(x)

print("**************PYTHON CLASSES**************\n")

print("\nejercicio 1:")
class MyClass:
  x="esto es una clase"
  print(x)
print("\nejercicio 2:")
class MyClass: 
  x= 5
p1=  MyClass() #objeto

print(p1.x)

print("\nejercicio 3:")

class Persona:
  def __init__(self,name,age) :
   self.name= name
   self.age = age
p1 = Persona("Juan", 30)
p2 = Persona("Maria", 25)
print(p1.name,p1.age)  # imprime "Juan"
print(p2.age)   # imprime 25

print("**************PYTHON INHERITANCE**************\n")

print("\nejercicio 1:")
class Student(Persona):
  def __init__(self,name,age,id) :
   super().__init__(name, age)
   self.id = id

s1 = Student("Juan", 20, "12345")
print(s1.name)  # imprime "Juan"
print(s1.age)   # imprime 20
print(s1.id)  # imprime "12345"   

print("**************PYTHON MODULES**************\n")
print("\nejercicio 1:")
print("import mymodule")
print("\nejercicio 2:")
print("import mymodule as mx")
print("\nejercicio 3:")
print("import mymodule")
print("print(dir(mymodule))")
print("\nejercicio 4:")
print("from mymodule import person1")