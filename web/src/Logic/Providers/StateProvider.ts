import { createContext, useContext } from "react";
import AppStore from "../State/AppStore";
import { FeedBackStore } from "../State/FeedBack";

interface IStoresContext {
  appStore: AppStore;
  feedBackStore : FeedBackStore
}

export const StoresContext = createContext<IStoresContext>(
  {} as IStoresContext
);

export const useStores = () => useContext(StoresContext);
