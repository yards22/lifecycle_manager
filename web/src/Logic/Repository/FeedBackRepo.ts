import { DummyTestDataFeedBack } from "../../Data/DummyFeedBackData";
import { MFeedBack } from "../Modal/MFeedBack";
import {Request} from "../Utils/Fetch"
import { CheckResponse } from "../Utils/ResponseHandler";
import { AuthHeaders, AuthHeadersWithoutToken } from "../Utils/AuthHeaders";

export class FeedBackRepo{
    rq : Request;
    baseUrl : string;

    constructor(baseURL : string,rq:Request){
        this.baseUrl = baseURL
        this.rq = rq
    }

    async getFeedBacks():Promise<MFeedBack[]| undefined>{
        try{
            return DummyTestDataFeedBack;
        }catch(e){

        }
    }
}