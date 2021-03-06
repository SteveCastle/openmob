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
  minHeight,
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
  ${minHeight}
  box-sizing: border-box;
  position: relative;
`;

function Row({ children, container, tracing, disableSpacing }) {
  return (
    <StyledRow
      display={'flex'}
      flexWrap={['wrap', 'nowrap']}
      border={0}
      mx={disableSpacing ? 0 : 2}
      my={disableSpacing ? 0 : 1}
      width={1}
      minHeight={100}
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
  tracing: PropTypes.number,
  disableSpacing: PropTypes.bool
};
Row.defaultProps = { container: false };

export default Row;
