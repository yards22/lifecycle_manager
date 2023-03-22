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
        this.token = window.localStorage.getItem("token");
    }

    @action
    SetIsUserLoggedIn = (b:boolean)=>{
        this.isUserLoggedIn = b
    }

    @action
    SetUserMailId = (mail:string)=>{
        this.userMailId = mail
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
    PerformLogout = async ()=>{
        try{
            const res = await this.authRepo.Logout(this.token||"")
            if( res!==401) {
                this.SetToken(null);
                this.SetIsUserLoggedIn(false);
            }
        }
        catch(e){
          console.log(e);
        }
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
            this.SetToken(response.data.token)
            this.SetIsUserLoggedIn(true)
            this.SetUserMailId(mail_id)
            return response.status
        }catch(e){
            console.log(e)
        }
    }

    @action
    CheckIfLogin=async ()=>{
         this.isLoading = true;
         try {
            const { user_data } = await this.authRepo.me(this.token || "");
            this.userMailId = user_data;
          } catch (err) {
            throw err;
          } finally {
            this.isLoading=false;
          }
    }
}