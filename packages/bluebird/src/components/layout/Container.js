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
  flexWrap,
  maxWidth
} from 'styled-system';

const StyledContainer = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  ${maxWidth}
`;

function Container({ dark, children, tracing }) {
  return (
    <StyledContainer
      display={'flex'}
      border={0}
      flexWrap="wrap"
      m="auto"
      width={1}
      maxWidth={7}
      bg={tracing ? `violet.${tracing}` : null}
      fontSize={1}
    >
      {children}
    </StyledContainer>
  );
}

Container.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Container.defaultProps = {};

export default Container;
