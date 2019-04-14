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

const StyledElection = styled.section`
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

function SimpleElection({ title = 'SimpleElection' }) {
  return (
    <StyledElection
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="red.8"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledElection>
  );
}

SimpleElection.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleElection.defaultProps = { title: 'Simple Election' };

export default SimpleElection;
