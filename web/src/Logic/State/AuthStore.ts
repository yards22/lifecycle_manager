import { action, makeAutoObservable, observable } from "mobx";
import { AuthRepo } from "../Repository/AuthRepo";

export class AuthStore{
    @observable isLoading: boolean = false;
    @observable token: string | null = null;
    @observable isUserLoggedIn : boolean = false
    authRepo : AuthRepo

    constructor(authRepo:AuthRepo){
        makeAutoObservable(this)
        this.authRepo = authRepo
    }

    @action
    SetIsUserLoggedIn = (b:boolean)=>{
        this.isUserLoggedIn = b
    }

    @action
    SendAuthOTP = async (mail_id:string)=>{
        try{
            const response = await this.authRepo.sendAuthOTP(mail_id) 
        }catch(e){

        }
    }
}