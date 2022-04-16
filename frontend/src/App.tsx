import { useEffect, useState } from "react";
import styled from "styled-components";
import { getList } from "./api";
import { Button } from "./components/button";
import { Container } from "./components/container";
import { Paragraph, Title } from "./components/typegraphics";

const Page = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const Row = styled.div`
  display: flex;
  justify-content: space-between;
  width: 100%;
  max-width: 50rem;
  padding: 0 10px;
  margin: 5px 0;
`;

function App() {
  const [list, setList] = useState<string[]>([]);

  useEffect(() => {
    loadList();
  }, []);

  const loadList = async () => {
    const list = await getList();
    setList(list);
  };

  const renderList = () => {
    return list.map((item) => {
      return (
        <Row>
          <Paragraph>{item}</Paragraph>
          <Button>Wake</Button>
        </Row>
      );
    });
  };

  return (
    <Page>
      <Container>
        <Title>Wake the Wizzard</Title>

        {renderList()}
      </Container>
    </Page>
  );
}

export default App;
