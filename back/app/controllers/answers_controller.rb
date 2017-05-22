class AnswersController < ApplicationController
  def index
    puts "index"
  end

  def create
    p @answer = Answer.new(answer_params)
    if @answer.save
      data = {
        "id":   @answer.id,
        "code": @answer.code,
        "lang": @answer.lang
      }
      # http://api:1323
      client = Faraday.new(:url => "http://localhost:4567")
      res = client.post do |req|
        req.url '/answer-data'
        req.headers['Content-Type'] = 'application/json'
        req.body = data.to_json
      end
      # return JSON.parse(res.body)["Result"]
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
