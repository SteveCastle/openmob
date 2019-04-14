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

const StyledBlogPost = styled.section`
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

function SimpleBlogPost({ title }) {
  return (
    <StyledBlogPost
      display={'flex'}
      justifyContent="center"
      alignItems="flex-end"
      bg="orange.8"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledBlogPost>
  );
}

SimpleBlogPost.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleBlogPost.defaultProps = { title: 'Simple Blogpost' };

export default SimpleBlogPost;
