export class Project {
    _id?:number;
    Name:string;
    Description:string;
    Repository:string;
    Created_at?:Date;
    Author?:string;

    //PEDING: Review optional attributes. They are not assigned when creating the project in the client side but in the server. And they need those to be rendered as
    // existing projects.

    constructor(Name: string, Description: string, Repository: string, Creation: Date, Author: string) {
        this.Name = Name;
        this.Description = Description;
        this.Repository = Repository;
        this.Created_at = Creation;
        this.Author = Author;
    }
}
