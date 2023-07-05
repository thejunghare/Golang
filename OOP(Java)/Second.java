class Second {
    public static void main(String args[]){
        // This is accesing the Main method
        Main obj = new Main();
        System.out.println(obj.x);

        // Accessing public method from Main.java
        obj.publicMethod();
    }
}