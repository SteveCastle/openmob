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

function Row({ dark, children, tracing }) {
  return (
    <StyledRow
      display={'flex'}
      flexWrap={['wrap', 'nowrap']}
      border={0}
      mx={2}
      my={1}
      width={1}
      bg={
        tracing
          ? `fuschia.${tracing}`
          : null
      }
      fontSize={1}
      color={dark ? 'lightText' : 'darkText'}
    >
      {children}
    </StyledRow>
  );
}

Row.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Row.defaultProps = {};

export default Row;
