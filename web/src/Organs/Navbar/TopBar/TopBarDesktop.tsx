import { Title, useMantineTheme, Text } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import { Home, Globe, Bell, User, Search, Edit3, BarChart2 } from "react-feather";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { useStores } from "../../../Logic/Providers/StateProvider";

const STopBarContainer = styled.div`
  position: fixed;
  top: 0px;
  left: 0px;
  right: 0px;
  border-bottom: 1px solid #eaeaea;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0px 25px;
  height: 60px;
  width: 100%;
`;

const STopBar = styled.a`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-decoration: none;
  margin: 0 10px;
  color: ${(p) => p.theme.color};
  cursor: pointer;
  transition: all 0.3s;
  :hover {
    color: #525252;
  }
`;

function TopBarDesktop() {
  const stores = useStores();
  const mantineTheme = useMantineTheme();
  const navigate = useNavigate();
  return (
    <Observer>
      {() => {
        const { appStore } = stores;
        return (
          <STopBarContainer>
            <Title color={"black"} order={2}>
             Admin 22 Yards
            </Title>
            <div
              style={{
                display: "flex",
                alignItems: "center",
              }}
            >
              <STopBar
                theme={{
                  color:
                    appStore.navigationState === 0
                      ? mantineTheme.colors[mantineTheme.primaryColor][7]
                      : "gray",
                }}
                onClick={() => {
                  navigate("/feedback");
                  appStore.setNavigationState(0);
                }}
              >
                <Edit3 size={"20"} />
                <Text size="xs">Feedbacks</Text>
              </STopBar>
              <STopBar
                theme={{
                  color:
                    appStore.navigationState === 1
                      ? mantineTheme.colors[mantineTheme.primaryColor][7]
                      : "gray",
                }}
                onClick={() => {
                  navigate("/polls");
                  appStore.setNavigationState(1);
                }}
              >
                <BarChart2 size={"20"} />
                <Text size="xs">Polls</Text>
              </STopBar>
              <STopBar
                theme={{
                  color:
                    appStore.navigationState === 2
                      ? mantineTheme.colors[mantineTheme.primaryColor][7]
                      : "gray",
                }}
                onClick={() => {
                  navigate("/stories");
                  appStore.setNavigationState(2);
                }}
              >
                <Bell size={"20"} />
                <Text size="xs">Stories</Text>
              </STopBar>
            </div>
          </STopBarContainer>
        );
      }}
    </Observer>
  );
}

export default TopBarDesktop;
