import { action, makeAutoObservable, observable } from "mobx";
import { AuthRepo } from "../Repository/AuthRepo";

const TOKEN_KEY = "token";

export class AuthStore{
    @observable isLoading: boolean = false;
    @observable token: string | null = null;
    @observable isUserLoggedIn : boolean = false
    @observable userMailId : string = ""
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
    SetToken = (token: string | null) => {
        let validToken: string | null = null;
        if (!token) validToken = null;
        else if (token !== "") validToken = token;
        if (validToken) window.localStorage.setItem(TOKEN_KEY, validToken);
        else window.localStorage.removeItem(TOKEN_KEY);
        this.token = validToken;
    };

    @action
    SendAuthOTP = async (mail_id:string)=>{
        try{
            const response = await this.authRepo.sendAuthOTP(mail_id) 
            return response
        }catch(e){
            console.log(e)
        }
    }

    @action
    VerifyAuthOTP = async (mail_id: string, OTP:string)=>{
        try{
            console.log(mail_id,OTP)
            const response = await this.authRepo.verifyAuthOTP(mail_id,OTP) 
            this.SetToken(response.data)
            return response.status
        }catch(e){
            console.log(e)
        }
    }
}