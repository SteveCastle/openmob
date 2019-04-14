import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap
} from 'styled-system';

import Container from './Container';

const StyledRow = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  box-sizing: border-box;

`;

function Row({ children, container, tracing }) {
  return (
    <StyledRow
      display={'flex'}
      flexWrap={['wrap', 'nowrap']}
      border={0}
      mx={2}
      my={1}
      width={1}
      bg={tracing ? `fuschia.${tracing}` : null}
      fontSize={1}
    >
      {container ? <Container>{children}</Container> : children}
    </StyledRow>
  );
}

Row.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  container: PropTypes.bool,
  tracing: PropTypes.number
};
Row.defaultProps = { container: false };

export default Row;
