import { AuthHeadersWithoutToken } from "../Utils/AuthHeaders";
import { Request } from "../Utils/Fetch";

export class AuthRepo{
    baseUrl : string;
    rq : Request

    constructor(baseUrl:string,rq:Request){
        this.baseUrl = baseUrl
        this.rq = rq
    }

    async sendAuthOTP(mail_id : string){
        try{
            const res = await this.rq.Post(this.baseUrl+"/sendOTP",{mail_id},AuthHeadersWithoutToken())
            return res.status
        }catch(e){
            console.log(e)
        }
    }

    async verifyAuthOTP(mail_id:string, OTP : string){
        try{
            const res = await this.rq.Post(this.baseUrl+"/login",{mail_id,OTP},AuthHeadersWithoutToken())
            const response = await res.json()
            return response
        }catch(e){
            console.log(e)
        }
    }

    async Logout(token:string){
        try{
           const res = await this.rq.Delete( `${this.baseUrl}/logout`,
           {},
           { Authorization: `Bearer ${token}` })
           return res.status
        }
        catch(e){
          console.log(e)
        }
    }

    async me(token:string){
        try{
            // will get mail id as resposne body
           const res = await this.rq.Post(this.baseUrl+"/me",token,AuthHeadersWithoutToken());
           const response = await res.json();
           return response
        }
        catch(err){
            console.log(err);
        }
    }
}