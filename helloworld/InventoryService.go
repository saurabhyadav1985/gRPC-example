import ballerina/grpc;
import ballerina/io;
endpoint grpc:Listener listener {
   host:"localhost",
   port:9000
};

@grpc:ServiceConfig
service InventoryService bind listener {
   getItemByName(endpoint caller, string value) {
       // Implementation goes here.
       // You should return a Items
   }

   getItemByID(endpoint caller, string value) {
       // Creating a dummy inventory item
       Item requested_item;
       requested_item.id = value;
       requested_item.name = "Sample Item " + value ;
       requested_item.description = "Sample Item Desc for " + value;
       _ = caller-&gt;send(requested_item);
       _ = caller-&gt;complete();
   }

   addItem(endpoint caller, Item value) {
       // Implementation goes here.
       // You should return a boolean
   }

}
