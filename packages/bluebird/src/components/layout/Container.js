import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledContainer = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Container({ dark, children, tracing }) {
  return (
    <StyledContainer
      border={0}
      p={3}
      bg={
        tracing
          ? `violet.${tracing}`
          : dark
          ? 'darkBackground'
          : 'lightBackground'
      }
      fontSize={1}
      color={dark ? 'lightText' : 'darkText'}
    >
      {children}
    </StyledContainer>
  );
}

Container.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Container.defaultProps = {};

export default Container;
