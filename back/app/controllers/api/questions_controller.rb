class Api::QuestionsController < ApplicationController
  def index
    @questions = Question.all
    render 'index', formats: 'json', handlers: 'jbuilder'
  end
end
