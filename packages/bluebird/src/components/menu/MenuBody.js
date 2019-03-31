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
  height,
  alignItems,
  justifyContent,
  flexWrap,
  flexDirection
} from 'styled-system';

const StyledMenuBody = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${alignItems}
  ${justifyContent}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  ${flexDirection}
  overflow: scroll;
  position: fixed;
  top: 13%;
  left: 0;
`;

function MenuBody({ children }) {
  return (
    <StyledMenuBody flexDirection="column" fontSize={4} height={'80%'}>
      {children}
    </StyledMenuBody>
  );
}

MenuBody.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ])
};

export default MenuBody;
