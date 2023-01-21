import { stringify } from "querystring";
import { MFeedBack } from "../Model/MFeedBack";
import {Request} from "../Utils/Fetch"
import { CheckResponse } from "../Utils/ResponseHandler";
import { AuthHeaders, AuthHeadersWithoutToken } from "../Utils/AuthHeaders";

export class FeedBackRepo{
    rq : Request;
    baseUrl : string;
    baseUrlForImages :string

    constructor(baseURL : string,baseImage:string,rq:Request){
        this.baseUrl = baseURL
        this.baseUrlForImages = baseImage
        this.rq = rq
    }

    async getFeedBacks(token:string){
        try{
            const res = await this.rq.Get(`${this.baseUrl}`,AuthHeaders(token));
            const {body} = await CheckResponse(res,201)
            const rawFeedBack = body.data as MFeedBack[]
            const finalFeedBacks: MFeedBack[] = [];
            rawFeedBack.forEach(feedBack=>{
                // feedBack.image_uri = this.baseUrlForImages+feedBack.image_uri;
                if (feedBack.image_uri !==null)
                feedBack.image_uri = (this.baseUrlForImages).concat(feedBack.image_uri as string)
                finalFeedBacks.push(feedBack)
            })
            return {
              feedback : finalFeedBacks
            };
        }
        catch(err){
           throw err
        }
    }

    async postFeedbackComments(status : boolean , comment : string | null,idx:number, token:string){
         try{
           const res = await this.rq.Post(`${this.baseUrl}/comment`,{status,comment,feedback_id:idx},AuthHeaders(token));
           const response = await res.json()
           return response
         }
         catch(err){
            throw err
         }
    }
}