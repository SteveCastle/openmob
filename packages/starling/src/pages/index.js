import React from 'react';
import PropTypes from 'prop-types';
import { graphql, Link } from 'gatsby';
import ThemeProvider from '@openmob/bluebird/src/ThemeProvider';
import skyward from '@openmob/bluebird/src/themes/skyward';
import ImageHero from '@openmob/bluebird/src/components/elements/heroes/ImageHero';
import Footer from '@openmob/bluebird/src/components/elements/footers/SimpleFooter';
import SubHeader from '@openmob/bluebird/src/components/type/SubHeader';
import ImageGrid from '@openmob/bluebird/src/components/lists/ImageGrid';
import ImageGridItem from '@openmob/bluebird/src/components/lists/ImageGridItem';
import Container from '@openmob/bluebird/src/components/layout/Container';

import Layout from '../components/Layout';
import SEO from '../components/SEO';
import { navigate } from '@reach/router';

const IndexPage = ({
  data: {
    wren: { listCause: causes = [] },
  },
}) => (
  <ThemeProvider theme={skyward}>
    <Layout title="grassroots.dev" id="List view" summary="Debug Mode">
      <SEO title="Home" keywords={[`gatsby`, `application`, `react`]} />
      <ImageHero title="Grassroots Dev" />
      <Container>
        <SubHeader>Featured Causes</SubHeader>
        <ImageGrid>
          {(causes || []).map(cause => (
            <ImageGridItem
              title={cause.Title}
              caption={cause.Summary}
              uri={cause.Photo.URI}
              width={4}
              onClick={() => navigate(`/${cause.Slug}`)}
            />
          ))}
        </ImageGrid>
      </Container>
      <Footer />
    </Layout>
  </ThemeProvider>
);

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      listCause(limit: 50) {
        ID
        Title
        Summary
        Slug
        Photo {
          URI
          Width
          Height
        }
      }
    }
  }
`;
IndexPage.propTypes = {
  data: PropTypes.shape({
    wren: PropTypes.shape({
      getCause: PropTypes.shape({}),
    }),
  }),
};

export default IndexPage;
