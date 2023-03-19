import { action, makeAutoObservable, observable } from "mobx";
import { MFeedBack } from "../Modal/MFeedBack";
import { FeedBackRepo } from "../Repository/FeedBackRepo";

export class FeedBackStore{
    @observable isLoading : boolean = false;
    @observable feedbackArray : MFeedBack[] = [];
    @observable token:string|null =null
    feedBackRepo : FeedBackRepo

    constructor(feedBackRepo:FeedBackRepo){
        makeAutoObservable(this)
        this.feedBackRepo = feedBackRepo
        this.token = window.localStorage.getItem("token");
    }

    @action
    SetFeedBack = (x:MFeedBack[])=>{
        this.feedbackArray = x;
    }

    @action
    SetIsLoading = (loading : boolean) =>{
        this.isLoading = loading
    }

    @action
    GetFeedBacks = async ()=> {
        this.SetIsLoading(true)
        try{
            let x = await this.feedBackRepo.getFeedBacks();
            if(x) this.SetFeedBack(x);
        }catch(e){
            throw e;
        }finally{
            this.SetIsLoading(false)
        }
    }
   }