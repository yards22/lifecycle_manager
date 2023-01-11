import { ActionIcon, useMantineTheme } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import { Edit3, BarChart2, Edit } from "react-feather";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { useStores } from "../../Logic/Providers/StateProvider";

const SBottomBar = styled.div`
  position: fixed;
  bottom: 0px;
  right: 0px;
  left: 0px;
  height: 60px;
  display: flex;
  z-index: 100;
  justify-content: space-around;
  align-items: center;
  box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px,
    rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px,
    rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
`;

function BottomBar() {
  const mantineTheme = useMantineTheme();
  const stores = useStores();
  const navigate = useNavigate();

  return (
    <Observer>
      {() => {
        const { appStore } = stores;
        return (
          <SBottomBar
            style={{
              // background: mantineTheme.colors["gray"][0],
              backgroundColor: "#ffffff",
              border: "1px solid " + mantineTheme.colors["gray"][3],
            }}
          >
            <ActionIcon
              color={
                appStore.navigationState === 0
                  ? mantineTheme.colors[mantineTheme.primaryColor][9]
                  : "gray"
              }
              onClick={() => {
                navigate("/feedback");
                appStore.setNavigationState(0);
              }}
            >
              <Edit3 size={"20"} />
            </ActionIcon>
            <ActionIcon
              color={
                appStore.navigationState === 1
                  ? mantineTheme.colors[mantineTheme.primaryColor][9]
                  : "gray"
              }
              onClick={() => {
                navigate("/polls");
                appStore.setNavigationState(1);
              }}
            >
              <BarChart2 size={"20"} />
            </ActionIcon>
            <ActionIcon
              color={
                appStore.navigationState === 2
                  ? mantineTheme.colors[mantineTheme.primaryColor][9]
                  : "gray"
              }
              onClick={() => {
                navigate("/stories");
                appStore.setNavigationState(2);
              }}
            >
              <Edit size={"20"} />
            </ActionIcon>
          </SBottomBar>
        );
      }}
    </Observer>
  );
}

export default BottomBar;
