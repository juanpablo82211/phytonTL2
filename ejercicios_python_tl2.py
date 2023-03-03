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

