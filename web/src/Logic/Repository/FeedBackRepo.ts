import { DummyTestDataFeedBack } from "../../Data/DummyFeedBackData";
import { MFeedBack } from "../Modal/MFeedBack";
import {Request} from "../Utils/Fetch"

export class FeedBackRepo{
    rq : Request;
    basUrl : string;
    
    constructor(baseURL : string,rq:Request){
        this.basUrl = baseURL
        this.rq = rq
    }

    async getFeedBacks():Promise<MFeedBack[]| undefined>{
        try{
            return DummyTestDataFeedBack;
        }catch(e){

        }
    }
}