import React, { useState } from 'react';

import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  fontSize,
  color,
  borders,
  borderRadius,
  display
} from 'styled-system';

const StyledOverlay = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${fontSize}
  ${color}
  ${display}
  align-items: flex-start;
  justify-content: flex-start;
  position: absolute;
  box-sizing: border-box;
  top: ${({ nestingLevel }) => nestingLevel * 32}px;
  left: ${({ nestingLevel }) => nestingLevel * 32}px;
  width: 100%;
  height: 80%;
  transition: background-color .2s ease-out;
  z-index: ${({ locked }) => (locked ? 999 : 998)};
`;

function Overlay({ children, locked, onClick, nestingLevel }) {
  return (
    <StyledOverlay
      border={0}
      locked={locked}
      onClick={onClick}
      display="flex"
      nestingLevel={nestingLevel}
    >
      {locked && children}
    </StyledOverlay>
  );
}

Overlay.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  locked: PropTypes.bool,
  onClick: PropTypes.func
};
Overlay.defaultProps = {};

export default Overlay;
