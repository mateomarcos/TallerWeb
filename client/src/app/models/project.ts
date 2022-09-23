
export class Project {
    _id?:number;
    Name:string;
    Description:string;
    Repository: string;
    Created_at: Date;
    Author: string;


    //HACEN FALTA LOS DEMAS ATRIBUTOS DE EL MODEL DE GOLANG?

    constructor(Name: string, Description: string, Repository: string, Author:string, Created_at:Date) {
        this.Name = Name;
        this.Description = Description;
        this.Repository = Repository;
        this.Created_at = Created_at;
        this.Author = Author;
    }
}