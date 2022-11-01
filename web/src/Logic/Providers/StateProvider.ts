import { createContext, useContext } from "react";
import AppStore from "../State/AppStore";

interface IStoresContext {
  appStore: AppStore;
}

export const StoresContext = createContext<IStoresContext>(
  {} as IStoresContext
);

export const useStores = () => useContext(StoresContext);
