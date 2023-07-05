// Class of name Main
public class Main{
    // x => attribute of main class
    int x = 5;

    String fname, lname;

    // final keyword => value of y can't be changed
    final int y = 5;

    // Methods (static method)
    static void staticMethod(){
        System.out.println("Hello from static method!");
    }

    // Public method
    public void publicMethod(){
        System.out.println("Hello from public method!");
    }


    // class constructor for Main class
    public Main(String lastName){
        fname =  "OOPS";
        lname = lastName;
    }

    // Private attribute
    private String name;

    // Setter for attribute name
    public void setName(String myName){
        this.name = myName;
    }

    // Getter for attribute name
    public String getName(){
        return name;
    }



    // main method/function
    public static void main(String args[]){
        // creating objects to access 'x' from Main class
        Main obj = new Main("Java");

        // Modify attribute value, 'x' will be 10 now.
        obj.x = 10;

        // Accessing the attributes with help of object => obj
        System.out.println(obj.x);
        System.out.println("First name " + obj.fname);
        System.out.println("Last name "+ obj.lname);
        
        // A Class can have multiple objects
        // Main obj1 = new Main();

        // Calling static method => staticMethod()
        staticMethod();

        // Calling public method => publicMethod()
        obj.publicMethod();

        // Accessing encapsulated data
        obj.setName("John Cena");
        System.out.println(obj.getName());
    }
}