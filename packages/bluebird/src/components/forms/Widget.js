import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  display,
  justifyContent,
  borderRadius,
} from 'styled-system';

const StyledWidget = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${display}
  ${justifyContent}
  ${fontSize}
  ${color}
`;

function Widget({ children, block }) {
  return (
    <StyledWidget
      border={0}
      p={3}
      fontSize={1}
      width={block ? 1 : null}
      display="flex"
      justifyContent="flex-end"
    >
      {children}
    </StyledWidget>
  );
}

Widget.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
  state: PropTypes.oneOf(['ready', 'valid', 'error']),
};
export default Widget;
