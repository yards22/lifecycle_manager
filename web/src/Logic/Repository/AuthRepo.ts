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
            const res = await this.rq.Post(this.baseUrl+"/sendOTP",mail_id,AuthHeadersWithoutToken())
            console.log(res)
        }catch(e){
            console.log(e)
        }
    }
}