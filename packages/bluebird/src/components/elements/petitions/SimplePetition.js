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

const StyledPetition = styled.section`
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

function SimplePetition({ title }) {
  return (
    <StyledPetition
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="yellow.4"
    >
      <Title color="gray.0" size={5}>
        {title}
      </Title>
    </StyledPetition>
  );
}

SimplePetition.propTypes = {
  title: PropTypes.string.isRequired
};
SimplePetition.defaultProps = { title: 'Simple Petition' };

export default SimplePetition;
