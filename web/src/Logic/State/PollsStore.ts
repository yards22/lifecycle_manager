import { action, makeAutoObservable, observable } from "mobx";
import { PollsRepo } from "../Repository/PollsRepo";

export class PollsStore{
    @observable isLoading:boolean = false
    @observable pollsArray : any[] = []
    pollsRepo:PollsRepo

    constructor(pollsRepo:PollsRepo){
        makeAutoObservable(this)
        this.pollsRepo = pollsRepo
    }

    @action
    GetPolls = async (token:string)=>{
        try {
            const response = await this.pollsRepo.getPolls(token);
            console.log("get polls response",response)
        } catch (error) {
            throw error
        }
    }

    @action
    AddPoll = async (pollQuestion:string, pollOptions:string[], token:string)=>{
        try {
            const response = await this.pollsRepo.addPoll(pollQuestion ,pollOptions ,token);
            console.log("add poll response",response)
        } catch (error) {
            throw error
        }
    }

    
}