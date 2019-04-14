import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledBoycott = styled.section`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  min-height: 400px;
  `;
const Title = styled.h1`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  `;

function SimpleBoycott({ title }) {
  return (
    <StyledBoycott
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="orange.8"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledBoycott>
  );
}

SimpleBoycott.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleBoycott.defaultProps = { title: 'Simple Boycott' };

export default SimpleBoycott;
