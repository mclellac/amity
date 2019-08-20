import React, { Component } from 'react';

const API = 'http://localhost:3001/v1/post/';
const QUERY = '1';

class Fetch extends Component {
  constructor(props) {
    super(props);

    this.state = {
      id: [],
      isLoading: false,
      error: null,
    };
  }

  componentDidMount() {
    this.setState({ isLoading: true });

    fetch(API + QUERY)
      .then(response => {
        if (response.ok) {
          return response.json();
        } else {
          throw new Error('Something went wrong ...');
        }
      })
      .then(data => this.setState({ 
        id: data.id,
        created: data.created,
        title: data.title,
        article: data.article,
        isLoading: false }))
      .catch(error => this.setState({ error, isLoading: false }));
  }

  render() {
    const { id, created, title, article, isLoading, error } = this.state;

    if (error) {
      return <p>{error.message}</p>;
    }

    if (isLoading) {
      return <p>Loading ...</p>;
    }

    return (
      <div class="post">
        <div>{id}</div>
        <div>{title}</div>
        <div> created on: {created}</div>
        <div>{article}</div>
      </div>
    );
  }
}

export default Fetch;
