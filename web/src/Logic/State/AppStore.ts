import { observable, action, makeAutoObservable } from "mobx";

export default class AppStore{
    @observable currentTab: number = 0;
    @observable currentWindowSize : number = 0;
    @observable selectedDate : Date = new Date();

    constructor(){
        makeAutoObservable(this)
    }

    @action
    SetCurrentTab = (currentTab:number)=>{
        this.currentTab = currentTab;
    }

    @action
    SetCurrentWindowSize = (size : number)=>{
        this.currentWindowSize = size
    }

    @action
    SetCurrentDate = (date : Date) =>{
        this.selectedDate = date
    }
}