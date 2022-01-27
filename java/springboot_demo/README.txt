1. run the application using
   ./mvnw spring-boot:run 
2. Or you can build the JAR file with
   ./mvnw clean package
   Then you can run the JAR file:
   java -jar target/springboot_demo-0.0.1-SNAPSHOT.jar
   
   if you want to assign port, use:
   java -jar target/springboot_demo-0.0.1-SNAPSHOT.jar --server.port=80