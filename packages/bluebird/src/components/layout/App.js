import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
  flexDirection
} from 'styled-system';

const StyledApp = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${flexWrap}
  ${flexDirection}
  ${color}
  ${display}
`;

function App({ children }) {
  return (
    <StyledApp
      display={'flex'}
      border={0}
      flexWrap="wrap"
      flexDirection="column"
      width={1}
      height={5}
      fontSize={1}
    >
      {children}
    </StyledApp>
  );
}

App.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
};
App.defaultProps = {};

export default App;
