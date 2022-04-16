import styled from "styled-components";
import { scaling } from "./_config";

export const Container = styled.div`
  @media (min-width: ${scaling.sm}) {
    max-width: ${scaling.sm};
  }

  @media (min-width: ${scaling.md}) {
    max-width: ${scaling.md};
  }

  @media (min-width: ${scaling.lg}) {
    max-width: ${scaling.lg};
  }

  @media (min-width: ${scaling.xl}) {
    max-width: ${scaling.xl};
  }

  margin: 10px;
  width: 100%;
`;
