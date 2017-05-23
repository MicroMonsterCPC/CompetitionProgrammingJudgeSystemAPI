require "faraday"
require "json"

class AnswersController < ApplicationController
  def index
    puts "index"
  end

  def create
    p @answer = Answer.new(answer_params)
    if @answer.save
      data = {
        "id":   @answer.question_id.to_s,
        "code": @answer.code,
        "lang": @answer.lang
      }
      client = Faraday.new(:url => "http://api:1323")
      res = client.post do |req|
        req.url '/answer-data'
        req.headers['Content-Type'] = 'application/json'
        req.body = data.to_json
      end
      redirect_to :success
    else
      puts "Error"
    end
  end
  private
  def answer_params
    params.require(:answer).permit(:question_id, :code, :lang)
  end
end
