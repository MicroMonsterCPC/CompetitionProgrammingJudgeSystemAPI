require "faraday"
require "json"

class AnswersController < ApplicationController
  protect_from_forgery :except => [:create]
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
      client = Faraday.new "http://api:1323" do |b|
        b.adapter Faraday.default_adapter
        b.basic_auth ENV{"AUTH_USERNAME"}, ENV["AUTH_PASSWORD"]
      end
      res = client.post do |req|
        req.url '/judgement-answer'
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
