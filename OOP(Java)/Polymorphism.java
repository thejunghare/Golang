class Animal {
    public void animalSound(){
        System.out.println("The animal makes sound");
    }
}

class Pig extends Animal{
    public void animalSound(){
        System.out.println("The pig says: wee wee");
    }
}

class Dog extends Animal{
    public void animalSound(){
        System.out.println("The dog says: bow bow");
    }
}

class Cat extends Animal{
    public void animalSound(){
        System.out.println("The cat says: meow meow");
    }
}

class Polymorphism {
    public static void main (String args[]){
        Animal obj = new Animal();
        Animal objPig = new Pig();
        Animal objDog = new Dog();
        Animal objCat = new Cat();

        obj.animalSound();
        objPig.animalSound();
        objDog.animalSound();
        objCat.animalSound();
    }
}