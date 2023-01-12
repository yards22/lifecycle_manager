import { action, makeAutoObservable, observable } from "mobx";
import { MFeedBack } from "../Modal/MFeedBack";
import { FeedBackRepo } from "../Repository/FeedBackRepo";

export class FeedBackStore{
    @observable isLoading : boolean = false
    @observable feedbackArray : MFeedBack[] = []
    feedBackRepo : FeedBackRepo

    constructor(feedBackRepo:FeedBackRepo){
        makeAutoObservable(this)
        this.feedBackRepo = feedBackRepo
    }

    @action
    SetIsLoading = (loading : boolean) =>{
        this.isLoading = loading
    }

    @action
    GetFeedBacks = async ()=> {
        try{

        }catch(e){

        }
    }
}